import { postRequest } from '../../repositories/postRequest';
type TParams = {
  mail: string;
  password: string;
};
export const useLogin = ({ mail, password }: TParams) => {
  postRequest<TParams>({
    url: 'user/login',
    formData: { mail, password },
  });
};
