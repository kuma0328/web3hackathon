import { useRouter } from 'next/router';
import { Board } from '../../../components/BaseParts/Board';
import { Community } from '../../../components/Community/Community';
import { Recipe } from '../../../components/Recipe/Recipe';
import { Page } from '../../../components/Wrapper/Page';
import { useGetCommunityById } from '../../../hooks/Community/useGetCommunityById';

const recipe = () => {
  const router = useRouter();
  const { communityId } = router.query;
  console.log(communityId);
  if (!communityId) {
    return <div>コミュニティIdの取得中</div>;
  }
  const { data, error, isLoading } = useGetCommunityById(communityId as string);
  if (isLoading) {
    return <div>ローディング中です</div>;
  }
  if (error) {
    return <div>エラーが起こりました</div>;
  }
  const mock = {
    name: '寿司コミュ',
    content: '日本の寿司のコミュニティを盛り上げます',
    image_url: '/images/mock/1.jpg',
  };
  return (
    <Page wide={true}>
      <Board className="w-full">
        <div className="grid grid-cols-1 lg:grid-cols-2">
          <Community
            title={mock.name}
            description={mock.content}
            image={mock.image_url}
            recipeLink={`/community/${communityId}/recipe`}
            onClickJoin={() => 1}
          />
          <Recipe />
        </div>
      </Board>
    </Page>
  );
};

export default recipe;
