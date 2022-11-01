import create from 'zustand';

type State = {
  materials: string[];
  processes: string[];
  setMaterial: (state: string) => void;
  setProcess: (state: string) => void;
  deleteMaterial: (state: string) => void;
  deleteProcess: (state: string) => void;
};

export const useRecipeStore = create<State>((set) => ({
  materials: [],
  processes: [],
  description: '',
  setMaterial: (material) =>
    set((state) => ({
      materials: [...state.materials, material],
    })),
  setProcess: (process) =>
    set((state) => ({
      processes: [...state.processes, process],
    })),
  deleteMaterial: (material) =>
    set((state) => ({
      materials: state.materials.filter((m) => {
        m !== material;
      }),
    })),
  deleteProcess: (process) =>
    set((state) => ({
      processes: state.processes.filter((m) => {
        m !== process;
      }),
    })),
}));
