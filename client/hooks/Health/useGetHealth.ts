import { useQuery } from 'react-query';
import { getRequest } from '../../repositories/getRequest';

export const useGetHealth = () => {
  const { data, error, isLoading } = useQuery<{ data: string }, Error>(
    'health',
    () => getRequest('health')
  );
  return { data, error, isLoading };
};
