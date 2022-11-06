import { useRouter } from 'next/router';
import { Board } from '../../../components/BaseParts/Board';
import { Chats } from '../../../components/Chats/Chats';
import { Community } from '../../../components/Community/Community';
import { Page } from '../../../components/Wrapper/Page';
import { useGetCommunityById } from '../../../hooks/Community/useGetCommunityById';
const communityDetail = () => {
  const router = useRouter();
  const { communityId } = router.query;
  if (!communityId) {
    return <div>コミュニティIdの取得中</div>;
  }
  // hooksのルールに反しているけど今のところ解決方法が思いつかない
  const { data, error, isLoading } = useGetCommunityById(
    (communityId as string) ?? ''
  );

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
    <Page
      wide={false}
      className="flex min-h-screen flex-col items-center pt-20 lg:px-5"
    >
      <div className="grid w-full grid-cols-1 gap-5 lg:grid-cols-2">
        <Board>
          <Community
            title={mock.name}
            description={mock.content}
            image={mock.image_url}
            recipeLink=""
            onClickJoin={() => 1}
          />
        </Board>
        <Chats />
      </div>
    </Page>
  );
};

export default communityDetail;
