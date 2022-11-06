import { Board } from '../../components/BaseParts/Board';
import { Page } from '../../components/Wrapper/Page';
import { Form } from '../../components/User/Form';

export const signup = () => {
  return (
    <Page wide={true}>
      <Board className="w-full md:w-1/2">
        <Form title="ユーザー登録" />
      </Board>
    </Page>
  );
};
export default signup;
