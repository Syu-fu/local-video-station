import { useEffect, useState, type FC } from 'react';
import { useParams } from 'react-router-dom';
import getVideoDetail from '../api/getVideoDetail';
import DrawerMenuComponent from '../components/DrawerMenuComponent';
import VideoEditForm from '../components/VideoEditForm';
import type Video from '../types/video';

const VideoEditPage: FC = () => {
  type Param = {
    id?: string;
  };
  const params: Param = useParams<Param>();
  const id = params?.id;

  const [video, setVideo] = useState<Video>();

  useEffect(() => {
    if (id === undefined) return;
    getVideoDetail(id)
      .then((data) => {
        setVideo(data);
      })
      .catch((error) => {
        console.error(error);
      });
  }, [id]);

  return (
    <>
      <DrawerMenuComponent />
      {video !== undefined && <VideoEditForm video={video} />}
    </>
  );
};

export default VideoEditPage;
