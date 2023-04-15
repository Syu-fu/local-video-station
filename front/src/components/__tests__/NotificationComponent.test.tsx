import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import NotificationComponent from '../NotificationComponent';

describe('Notification', () => {
  test('renders success notification with messages', () => {
    const handleClose = vitest.fn();

    render(
      <NotificationComponent
        open={true}
        severity="success"
        messages={['Success message 1', 'Success message 2']}
        onClose={handleClose}
      />
    );

    expect(screen.getByText('Success message 1')).toBeInTheDocument();
    expect(screen.getByText('Success message 2')).toBeInTheDocument();
  });

  test('renders error notification with messages', () => {
    const handleClose = vitest.fn();

    render(
      <NotificationComponent
        open={true}
        severity="error"
        messages={['Error message 1', 'Error message 2']}
        onClose={handleClose}
      />
    );

    expect(screen.getByText('Error message 1')).toBeInTheDocument();
    expect(screen.getByText('Error message 2')).toBeInTheDocument();
  });

  test('calls onClose when clicking the close button', () => {
    const handleClose = vitest.fn();

    render(
      <NotificationComponent
        open={true}
        severity="success"
        messages={['Success message 1']}
        onClose={handleClose}
      />
    );

    fireEvent.click(screen.getByRole('button'));
    expect(handleClose).toHaveBeenCalledTimes(1);
  });
});
