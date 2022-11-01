import create from 'zustand';

type State = {
  chatImage: string;
  chatContent: string;
  setChatImage: (state: string) => void;
  setChatContent: (state: string) => void;
};

export const useChatContentStore = create<State>((set) => ({
  chatImage: '',
  chatContent: '',
  setChatImage: (image) =>
    set((state) => ({
      chatImage: image,
    })),
  setChatContent: (content) =>
    set(() => ({
      chatContent: content,
    })),
}));
