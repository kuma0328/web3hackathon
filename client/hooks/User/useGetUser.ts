import { useQuery } from 'react-query';
import { getRequest } from '../../repositories/getRequest';
type TResponse = {
  id: number;
  name: string;
  mail: string;
  password: string;
};
export const useGetUser = (id: number) => {
  const { data, error, isLoading } = useQuery<{ data: TResponse }, Error>(
    'user',
    () => getRequest(`user/${id}`)
  );
  return { data, error, isLoading };
};
