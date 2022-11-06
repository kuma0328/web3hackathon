import create from 'zustand';

type State = {
  image: FileList | null;
  title: string;
  description: string;
  setImage: (state: FileList) => void;
  setTitle: (state: string) => void;
  setDescription: (state: string) => void;
};

export const useCommunityStore = create<State>((set) => ({
  image: null,
  title: '',
  description: '',
  setImage: (image) =>
    set((state) => ({
      image: image,
    })),
  setTitle: (title) =>
    set(() => ({
      title: title,
    })),
  setDescription: (description) =>
    set(() => ({
      description: description,
    })),
}));
