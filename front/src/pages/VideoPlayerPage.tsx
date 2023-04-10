import { type FC, useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import getVideoDetail from '../api/getVideoDetail';
import VideoDetailComponent from '../components/VideoDetailComponent';
import VideoPlayerComponent from '../components/VideoPlayerComponent';
import type Video from '../types/video';

const VideoPlayerPage: FC = () => {
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
      {video !== undefined && <VideoDetailComponent video={video} />}
      {video !== undefined && <VideoPlayerComponent url={video.url} />}
    </>
  );
};

export default VideoPlayerPage;
