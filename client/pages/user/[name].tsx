import { Board } from '../../components/BaseParts/Board';
import { TypoGraphy } from '../../components/BaseParts/TypoGraphy';
import { Page } from '../../components/Wrapper/Page';
import { useRouter } from 'next/router';
import { Descriptions } from '../../components/BaseParts/Descriptions';
import { BelongingCommunity } from '../../components/Community/BelongingCommunity';
import { Button } from '../../components/BaseParts/Button';
import { useLogOut } from '../../hooks/User/useLogOut';

const index = () => {
  const router = useRouter();
  const { name } = router.query;
  return (
    <div>
      <Page wide={true}>
        <Board className="w-full md:w-1/2">
          <div className="p-5 text-center md:p-20">
            <div className="m-auto my-3 h-32 w-32 rounded-full bg-orange-400"></div>
            <TypoGraphy className="text-2xl">{name}</TypoGraphy>
            <Descriptions>
              <p>Eth:0.99</p>
              <p className="py-1">いいね数:10</p>
              <p>コメント数:7</p>
            </Descriptions>
            <TypoGraphy className="text-left">所属コミュニティ</TypoGraphy>
            <BelongingCommunity className="rounded-md" />
            <div className="my-5">
              <Button onClick={useLogOut}>ログアウトする</Button>
            </div>
          </div>
        </Board>
      </Page>
    </div>
  );
};
export default index;
