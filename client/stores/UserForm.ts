import create from 'zustand';

type State = {
  name: string;
  mail: string;
  password: string;
  passwordConfirmation: string;
  setName: (state: string) => void;
  setMail: (state: string) => void;
  setPassword: (state: string) => void;
  setPasswordConfirmation: (state: string) => void;
};

export const UserForm = create<State>((set) => ({
  name: '',
  mail: '',
  password: '',
  passwordConfirmation: '',
  setName: (name) =>
    set(() => ({
      name: name,
    })),
  setMail: (mail) =>
    set(() => ({
      mail: mail,
    })),
  setPassword: (password) =>
    set(() => ({
      password: password,
    })),
  setPasswordConfirmation: (pass) =>
    set(() => ({
      passwordConfirmation: pass,
    })),
}));
