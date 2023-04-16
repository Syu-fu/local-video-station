import { type FC } from 'react';
import DrawerMenuComponent from '../components/DrawerMenuComponent';
import VideoCreateForm from '../components/VideoCreateForm';

const VideoCreatePage: FC = () => {
  return (
    <>
      <DrawerMenuComponent />
      <VideoCreateForm />
    </>
  );
};

export default VideoCreatePage;
