import { useQuery } from 'react-query';
import { getRequest } from '../../repositories/getRequest';
type TResponse = {
  id: number;
  name: string;
  content: string;
  image_url: string;
  rrecipe: string | null;
  user: string | null;
};
export const useGetCommunityById = (id: string) => {
  const { data, error, isLoading } = useQuery<{ data: TResponse }, Error>(
    'communities',
    () => getRequest(`community/${id}`)
  );
  return { data, error, isLoading };
};
