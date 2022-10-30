import { TypoGraphy } from '../../components/BaseParts/TypoGraphy';
import { BelongingCommunity } from '../../components/Community/BelongingCommunity';
import { Communities } from '../../components/Community/Communities';
import { Page } from '../../components/Wrapper/Page';

const index = () => {
  return (
    <Page wide={false} className="p-10">
      <TypoGraphy>所属しているコミュニティ</TypoGraphy>
      <BelongingCommunity />
      <TypoGraphy>コミュニティ一覧</TypoGraphy>
      <Communities />
    </Page>
  );
};

export default index;
