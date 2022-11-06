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
  const mock = [
    {
      name: '寿司コミュ',
      content: '日本の寿司のコミュニティを盛り上げます',
      image_url: '/images/mock/1.jpg',
    },
    {
      name: 'もつ鍋',
      content: 'もつ鍋のいろいろな食べ方を共有し合います',
      image_url: '/images/mock/2.jpg',
    },
    {
      name: '海外郷土料理コミュ',
      content: '海外のニッチな郷土料理を紹介し合います',
      image_url: '/images/mock/3.jpg',
    },
  ];
  return (
    <div className="grid grid-cols-1 gap-5 lg:grid-cols-3 xl:grid-cols-4">
      {mock.map((community, i) => {
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
