import { type FC } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import VideoPlayerPage from '../pages/VideoPlayerPage';
import VideoListPage from '../pages/list';

const AppRouter: FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/list" element={<VideoListPage />} />
        <Route path="video">
          <Route path=":id" element={<VideoPlayerPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
};

export default AppRouter;
