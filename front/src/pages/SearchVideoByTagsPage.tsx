import { type FC, useState, useCallback, useEffect } from 'react';
import { Box } from '@mui/material';
import { useNavigate, useSearchParams } from 'react-router-dom';
import getTagList from '../api/getTagList';
import getVideoCountByTags from '../api/getVideoCountByTags';
import getVideoListByTags from '../api/getVideoListByTags';
import DrawerMenuComponent from '../components/DrawerMenuComponent';
import PaginationComponent from '../components/PaginationComponent';
import TagAutocompleteComponent from '../components/TagAutoCompleteComponent';
import VideoListComponent from '../components/VideoListComponent';
import type Tag from '../types/tag';
import type Video from '../types/video';

const SearchVideoByTagsPage: FC = () => {
  const [searchParams] = useSearchParams();
  const [tags, setTags] = useState<Tag[]>([]);
  const [selectedTags, setSelectedTags] = useState<Tag[]>([]);
  const [videos, setVideos] = useState<Video[]>([]);
  const [noResults, setNoResults] = useState<boolean>(false);
  const [pageCount, setPageCount] = useState<number>(1);
  const navigate = useNavigate();

  const pageQuery = searchParams.get('page') ?? '1';
  const currentPage = parseInt(pageQuery) ?? 1;

  useEffect(() => {
    getTagList()
      .then((fetchedTags) => {
        setTags(fetchedTags);
      })
      .catch((error) => {
        console.error('Error fetching tags:', error);
      });
  }, []);

  const fetchVideos = useCallback((selectedTags: Tag[], page: number) => {
    getVideoListByTags(selectedTags, page)
      .then((fetchedVideos) => {
        setVideos(fetchedVideos);
        setNoResults(false);
      })
      .catch((error: { response: { status: number } }) => {
        if (error.response.status === 404) {
          setNoResults(true);
          setVideos([]);
        } else {
          console.error('Error fetching videos:', error);
        }
      });

    getVideoCountByTags(selectedTags)
      .then((videoCount) => {
        setPageCount(Math.ceil(videoCount.count / 10));
      })
      .catch((error) => {
        console.error('Error fetching video count:', error);
      });
  }, []);

  const handleTagChange = (tags: Tag[]) => {
    setSelectedTags(tags);
    fetchVideos(tags, currentPage);
    navigate(`?page=1`);
  };

  useEffect(() => {
    fetchVideos(selectedTags, currentPage);
  }, [selectedTags, currentPage, fetchVideos]);

  return (
    <Box display="flex" flexDirection="column" minHeight="100vh">
      <DrawerMenuComponent />
      <Box
        sx={{
          position: 'sticky',
          top: 0,
          backgroundColor: 'background.default',
          zIndex: 2,
          width: '100%',
          padding: '8px',
        }}
      >
        <Box
          mx="auto"
          width="100%"
          minHeight="48px"
          display="flex"
          alignItems="center"
        >
          <TagAutocompleteComponent
            tags={tags}
            selectedTags={selectedTags}
            onChange={handleTagChange}
          />
        </Box>
      </Box>
      <Box
        flexGrow={1}
        display="flex"
        flexDirection="column"
        alignItems="center"
        width="100vw"
      >
        {noResults ? (
          <p style={{ textAlign: 'center', marginTop: '16px' }}>
            No videos found for the selected tags.
          </p>
        ) : (
          <VideoListComponent videos={videos} />
        )}
      </Box>
      <PaginationComponent pageCount={pageCount} currentPage={currentPage} />
    </Box>
  );
};

export default SearchVideoByTagsPage;
