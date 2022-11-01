import create from 'zustand';

type State = {
  material: string;
  process: string;
  materials: string[];
  processes: string[];
  changeMaterial: (state: string) => void;
  changeProcess: (state: string) => void;
  setMaterial: (state: string) => void;
  setProcess: (state: string) => void;
  deleteMaterial: (state: string) => void;
  deleteProcess: (state: string) => void;
};

export const useRecipeStore = create<State>((set) => ({
  material: '',
  process: '',
  materials: [],
  processes: [],
  changeMaterial: (material) =>
    set(() => ({
      material: material,
    })),
  changeProcess: (process) =>
    set(() => ({
      process: process,
    })),
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
      materials: state.materials.filter((m) => m !== material),
    })),
  deleteProcess: (process) =>
    set((state) => ({
      processes: state.processes.filter((p) => p !== process),
    })),
}));
