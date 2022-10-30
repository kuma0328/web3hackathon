import { Board } from '../../components/BaseParts/Board';
import { Page } from '../../components/Wrapper/Page';
import { Form } from '../../components/User/Form';
export const signin = () => {
  return (
    <Page>
      <Board className="w-full md:w-1/2">
        <Form
          title="ログイン"
          mail={''}
          password={''}
          onChangeMail={() => {
            console.log('aa');
          }}
          onChangePassword={() => {
            console.log('aa');
          }}
          onSubmit={() => {
            console.log('aa');
          }}
        />
      </Board>
    </Page>
  );
};
export default signin;
