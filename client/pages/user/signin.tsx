import { Board } from '../../components/BaseParts/Board';
import { Page } from '../../components/Wrapper/Page';
import { Form } from '../../components/User/Form';
export const signin = () => {
  return (
    <Page wide={true}>
      <Board className="w-full md:w-1/2">
        <Form title="ログイン" />
      </Board>
    </Page>
  );
};
export default signin;
