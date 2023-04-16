import { type FC, useState, Fragment } from 'react';
import ListIcon from '@mui/icons-material/List';
import LocalOfferIcon from '@mui/icons-material/LocalOffer';
import MenuIcon from '@mui/icons-material/Menu';
import VideoLibraryIcon from '@mui/icons-material/VideoLibrary';
import {
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  ListSubheader,
  IconButton,
  AppBar,
  Toolbar,
  Typography,
  ButtonBase,
} from '@mui/material';
import { Link } from 'react-router-dom';

type MenuItem = {
  id: string;
  label: string;
  icon: React.ReactElement;
  link: string;
};

type MenuGroup = {
  id: string;
  label: string;
  items: MenuItem[];
};

const menuGroups: MenuGroup[] = [
  {
    id: 'display',
    label: 'Display',
    items: [
      {
        id: 'list',
        label: 'List',
        icon: <ListIcon />,
        link: '/list',
      },
    ],
  },
  {
    id: 'video',
    label: 'Video',
    items: [
      {
        id: 'create-video',
        label: 'Register',
        icon: <VideoLibraryIcon />,
        link: '/video/create',
      },
    ],
  },
  {
    id: 'tag',
    label: 'Tag',
    items: [
      {
        id: 'create-tag',
        label: 'Register',
        icon: <LocalOfferIcon />,
        link: '/tag/create',
      },
    ],
  },
];

const DrawerMenu: FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleDrawer = () => {
    setIsOpen(!isOpen);
  };

  return (
    <>
      <AppBar position="fixed">
        <Toolbar>
          <IconButton
            edge="start"
            color="inherit"
            aria-label="menu"
            onClick={toggleDrawer}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" component="div">
            Local Video Station
          </Typography>
        </Toolbar>
      </AppBar>
      <Toolbar />
      <Drawer
        open={isOpen}
        onClose={toggleDrawer}
        variant="temporary"
        anchor="left"
        sx={{
          width: '300px',
          '.MuiDrawer-paper': {
            width: '300px',
          },
        }}
      >
        <List>
          {menuGroups.map((group) => (
            <Fragment key={group.id}>
              <ListSubheader>{group.label}</ListSubheader>
              {group.items.map((item) => {
                  return (
                      <ButtonBase
                          key={item.id}
                          component={Link}
                          to={item.link}
                          onClick={toggleDrawer}
                          sx={{ width: "100%" }}
                      >
                          <ListItem>
                              <ListItemIcon>{item.icon}</ListItemIcon>
                              <ListItemText primary={item.label} />
                          </ListItem>
                      </ButtonBase>
                  );
              })}
            </Fragment>
          ))}
        </List>
      </Drawer>
    </>
  );
};

export default DrawerMenu;
