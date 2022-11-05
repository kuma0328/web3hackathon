import { TypoGraphy } from '../../components/BaseParts/TypoGraphy';
import { BelongingCommunity } from '../../components/Community/BelongingCommunity';
import { Communities } from '../../components/Community/Communities';
import { Page } from '../../components/Wrapper/Page';

const index = () => {
  return (
    <Page wide={false} className="p-5 md:pt-20">
      <TypoGraphy className="my-10 text-2xl">
        所属しているコミュニティ
      </TypoGraphy>
      <BelongingCommunity className="rounded-full" />
      <TypoGraphy className="my-10 text-2xl">コミュニティ一覧</TypoGraphy>
      <Communities />
    </Page>
  );
};

export default index;
