import { type FC, useState, useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import getVideoCount from '../api/getVideoCount';
import getVideoList from '../api/getVideoList';
import PageNationComponent from '../components/PaginationComponent';
import VideoListComponent from '../components/VideoListComponent';
import type Video from '../types/video';

const VideoListPage: FC = () => {
  const [searchParams] = useSearchParams();
  const numberPerPage = 10;

  const pageQuery = searchParams.get('page') ?? '1';
  const currentPage = parseInt(pageQuery) ?? 1;

  const [videos, setVideos] = useState<Video[]>([]);
  const [videoCount, setVideoCount] = useState(0);

  useEffect(() => {
    getVideoList(pageQuery)
      .then((data) => {
        setVideos(data);
      })
      .catch((error) => {
        console.error(error);
      });
    getVideoCount()
      .then((data) => {
        setVideoCount(Math.ceil(data.count / numberPerPage));
      })
      .catch((error) => {
        console.error(error);
      });
  }, [pageQuery]);

  return (
    <>
      <VideoListComponent videos={videos} />
      <PageNationComponent pageCount={videoCount} currentPage={currentPage} />
    </>
  );
};

export default VideoListPage;
