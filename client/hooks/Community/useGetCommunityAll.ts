import { useQuery } from 'react-query';
import { getRequest } from '../../repositories/getRequest';
type TResponse = {
  id: number;
  name: string;
  mail: string;
  password: string;
};
export const useGetCommunityAll = () => {
  const { data, error, isLoading } = useQuery<{ data: TResponse }, Error>(
    'user',
    () => getRequest(`community/all`)
  );
  return { data, error, isLoading };
};
