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
export const useGetCommunityAll = () => {
  const { data, error, isLoading } = useQuery<{ data: TResponse[] }, Error>(
    'communities',
    () => getRequest(`community/all`)
  );
  return { data, error, isLoading };
};
