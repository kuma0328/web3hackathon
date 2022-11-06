import { getRequest } from '../../repositories/getRequest';
export const useLogOut = () => {
  getRequest(`user/logout`);
};
