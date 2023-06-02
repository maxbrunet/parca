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

import {ReactNode, createContext, useContext, useEffect, useMemo, useState} from 'react';

export interface KeyDownState {
  isShiftDown: boolean;
  keys: string[];
}

const DEFAULT_VALUE = {
  isShiftDown: false,
  keys: [],
};

const KeyDownContext = createContext<KeyDownState>(DEFAULT_VALUE);

export const KeyDownProvider = ({
  children,
}: {
  children: ReactNode;
  value?: KeyDownState;
}): JSX.Element => {
  const [keys, setKeys] = useState<string[]>([]);
  // Shift requires special handling because it is a shortcut key
  const isShiftDown = keys.includes('Shift') && !keys.includes('Meta');

  useEffect(() => {
    if (typeof window === 'undefined') {
      return;
    }

    const handleKeyDown = (event: {key: string}): void => {
      setKeys(keys => [...keys, event.key]);
    };

    window.addEventListener('keydown', handleKeyDown);

    const handleKeyUp = (event: {key: string}): void => {
      setKeys(keys => [...keys.filter(key => key !== event.key)]);
    };

    window.addEventListener('keyup', handleKeyUp);

    return () => {
      window.removeEventListener('keydown', handleKeyDown);
      window.removeEventListener('keyup', handleKeyUp);
    };
  }, []);

  const value = useMemo(() => ({keys, isShiftDown}), [keys, isShiftDown]);

  return <KeyDownContext.Provider value={value}>{children}</KeyDownContext.Provider>;
};

export const useKeyDown = (): KeyDownState => {
  const context = useContext(KeyDownContext);
  if (context == null) {
    return DEFAULT_VALUE;
  }
  return context;
};

export default KeyDownContext;
