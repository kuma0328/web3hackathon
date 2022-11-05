import { useGetCommunityAll } from '../../hooks/Community/useGetCommunityAll';
import { Board } from '../BaseParts/Board';
import { Community } from './Community';

export const Communities = () => {
  const { data, isLoading, error } = useGetCommunityAll();
  if (isLoading) {
    return <div>ローディング中</div>;
  }
  if (error) {
    return <div>なにかしらの不具合が発生しました</div>;
  }
  return (
    <div className="grid grid-cols-1 gap-5 lg:grid-cols-3 xl:grid-cols-4">
      {data?.data.map((community, i) => {
        return (
          <Board>
            <Community
              title={community.name}
              description={community.content}
              link={`/community/${community.name}`}
              image={community.image_url}
              recipeLink=""
              key={i}
            />
          </Board>
        );
      })}
    </div>
  );
};
