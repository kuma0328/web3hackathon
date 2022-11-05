import { postRequest } from '../../repositories/postRequest';
type TParams = {
  name: string;
  mail: string;
  password: string;
};
export const useSignUp = ({ name, mail, password }: TParams) => {
  postRequest<TParams, TParams>({
    url: 'user/signup',
    formData: { name, mail, password },
  });
};
