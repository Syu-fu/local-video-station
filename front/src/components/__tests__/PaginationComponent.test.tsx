import { render, screen } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';
import PaginationComponent from '../PaginationComponent';

describe('PaginationComponent', () => {
  const props = {
    pageCount: 5,
    currentPage: 2,
  };

  it('should render the pagination with correct props', () => {
    render(
      <Router>
        <PaginationComponent {...props} />
      </Router>
    );
    const pagination = screen.getByRole('navigation');
    expect(pagination).toBeInTheDocument();
    expect(pagination).toHaveAttribute('aria-label', 'pagination navigation');
    expect(screen.getByText('2')).toHaveClass('Mui-selected');
  });
});
