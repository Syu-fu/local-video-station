import { type FC } from 'react';
import DrawerMenuComponent from '../components/DrawerMenuComponent';
import TagCreateForm from '../components/TagCreateForm';

const TagCreatePage: FC = () => {
  return (
    <>
      <DrawerMenuComponent />
      <TagCreateForm />
    </>
  );
};

export default TagCreatePage;
