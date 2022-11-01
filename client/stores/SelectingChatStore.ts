import create from 'zustand';

type State = {
  recipe: boolean;
  random: boolean;
  tips: boolean;
  setRecipeChat: () => void;
  setRandomChat: () => void;
  setTipsChat: () => void;
};

export const useSelectingChatStore = create<State>((set) => ({
  recipe: true,
  random: false,
  tips: false,
  setRecipeChat: () =>
    set(() => ({
      recipe: true,
      random: false,
      tips: false,
    })),
  setRandomChat: () =>
    set(() => ({
      recipe: false,
      random: true,
      tips: false,
    })),
  setTipsChat: () =>
    set(() => ({
      recipe: false,
      random: false,
      tips: true,
    })),
}));
