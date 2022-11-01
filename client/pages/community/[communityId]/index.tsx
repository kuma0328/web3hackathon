import { useRouter } from 'next/router';
import { Board } from '../../../components/BaseParts/Board';
import { Chats } from '../../../components/Chats/Chats';
import { Community } from '../../../components/Community/Community';
import { Page } from '../../../components/Wrapper/Page';
const communityDetail = () => {
  const router = useRouter();
  const { communityId } = router.query;
  return (
    <Page
      wide={false}
      className="flex min-h-screen flex-col items-center pt-20 lg:px-5"
    >
      <div className="grid w-full grid-cols-1 gap-5 lg:grid-cols-2">
        <Board>
          <Community
            title="焼肉同好会"
            description="焼肉同好会では各地のさまざまな焼き肉を食べながら交流を深めて行きます"
            recipeLink="a"
            onClickJoin={() => 1}
          />
        </Board>
        <Chats />
      </div>
    </Page>
  );
};

export default communityDetail;
