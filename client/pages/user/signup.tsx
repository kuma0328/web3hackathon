import { Board } from '../../components/BaseParts/Board';
import { Page } from '../../components/Wrapper/Page';
import { Form } from '../../components/User/Form';

export const signup = () => {
  return (
    <Page wide={true}>
      <Board className="w-full md:w-1/2">
        <Form
          title="ユーザー登録"
          mail={''}
          password={''}
          name={''}
          passwordConfirmation={''}
          onChangeName={() => {
            console.log('aa');
          }}
          onChangeMail={() => {
            console.log('aa');
          }}
          onChangePassword={() => {
            console.log('aa');
          }}
          onChangePasswordConfirmation={() => {
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
export default signup;
