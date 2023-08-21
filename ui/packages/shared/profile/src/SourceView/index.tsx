// Copyright 2022 The Parca Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React, {useMemo} from 'react';

import {tableFromIPC} from 'apache-arrow';
import hljs from 'highlight.js';

import {Source} from '@parca/client';
import {useURLState} from '@parca/components';

import {LineNo} from './LineNo';

interface SourceViewProps {
  loading: boolean;
  data?: Source;
  total: bigint;
  filtered: bigint;
}

interface HighlighterProps {
  content: string;
  language?: string;
}

function Highlighter({content, language}: HighlighterProps): JSX.Element {
  const highlighted = useMemo(() => {
    return language !== undefined && language !== null
      ? hljs.highlight(language, content)
      : hljs.highlightAuto(content);
  }, [content, language]);

  return (
    <pre className="hljs">
      <code dangerouslySetInnerHTML={{__html: highlighted.value}} />
    </pre>
  );
}

export const SourceView = React.memo(function SourceView({
  data,
  loading,
  total,
  filtered,
}: SourceViewProps): JSX.Element {
  const [sourceLine] = useURLState({param: 'source_line', navigateTo: () => {}});
  if (loading || data === undefined) return <>Profile has no samples</>;

  const table = tableFromIPC(data.record);
  const cumulative = table.getChild('cumulative');
  const flat = table.getChild('flat');
  const lines = Array.from({length: flat?.length ?? 0}, (_, i) => i + 1);

  let cumulativeValues = '';
  for (let i = -1, n = cumulative?.length ?? 0; ++i < n; ) {
    const row = cumulative?.get(i) ?? 0;
    if (row > 0) {
      if (filtered > 0) {
        const unfilteredPercent = ((Number(row) / Number(total + filtered)) * 100).toFixed(2);
        const filteredPercent = ((Number(row) / Number(total)) * 100).toFixed(2);
        cumulativeValues +=
          row.toString() + '(' + unfilteredPercent + '% / ' + filteredPercent + '%)\n';
      } else {
        const percent = ((Number(row) / Number(total)) * 100).toFixed(2);
        cumulativeValues += row.toString() + '(' + percent + '%)\n';
      }
    } else {
      cumulativeValues += '\n';
    }
  }

  let flatValues = '';
  for (let i = -1, n = flat?.length ?? 0; ++i < n; ) {
    const row = flat?.get(i) ?? 0;
    if (row > 0) {
      if (filtered > 0) {
        const unfilteredPercent = ((Number(row) / Number(total + filtered)) * 100).toFixed(2);
        const filteredPercent = ((Number(row) / Number(total)) * 100).toFixed(2);
        flatValues += row.toString() + '(' + unfilteredPercent + '% / ' + filteredPercent + '%)\n';
      } else {
        const percent = ((Number(row) / Number(total)) * 100).toFixed(2);
        flatValues += row.toString() + '(' + percent + '%)\n';
      }
    } else {
      flatValues += '\n';
    }
  }

  return (
    <div className="h-[80vh] overflow-y-auto text-xs dark:bg-[#22272e] dark:text-[#adbac7]">
      <div className="flex gap-2">
        <pre>
          <code>
            Line{'\n'}
            {lines.map(line => (
              <LineNo value={line} key={line} isCurrent={sourceLine === line.toString()} />
            ))}
          </code>
        </pre>
        <pre>
          <code>{'Cummulative\n' + cumulativeValues}</code>
        </pre>
        <pre>
          <code>{'Flat\n' + flatValues}</code>
        </pre>
        <div className="overflow-x-auto">
          <pre>{'Source Code\n'}</pre>
          <Highlighter content={data.source} />
        </div>
      </div>
    </div>
  );
});

export default SourceView;
