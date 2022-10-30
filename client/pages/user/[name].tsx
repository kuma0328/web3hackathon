import { Board } from '../../components/BaseParts/Board';
import { TypoGraphy } from '../../components/BaseParts/TypoGraphy';
import { Page } from '../../components/Wrapper/Page';
import { useRouter } from 'next/router';
const index = () => {
  const router = useRouter();
  const { name } = router.query;
  return (
    <div>
      <Page wide={true}>
        <Board className="w-full md:w-1/2">
          <div className="p-5 text-center">
            <div className="m-auto my-3 h-32 w-32 rounded-full bg-orange-400"></div>
            <TypoGraphy className="text-xl">{name}</TypoGraphy>
            <div className="my-5 border-t-2 border-b-2 border-black py-5">
              <p>Eth:0.99</p>
              <p>いいね数:10</p>
              <p>コメント数:7</p>
            </div>
            <TypoGraphy className="text-left">所属コミュニティ</TypoGraphy>
            <div className="grid grid-cols-3 justify-items-center gap-5">
              {[0, 1, 2, 3].map((index) => {
                return (
                  <div
                    key={index}
                    className=" h-20 w-full rounded-md bg-orange-400"
                  ></div>
                );
              })}
            </div>
          </div>
        </Board>
      </Page>
    </div>
  );
};
export default index;
