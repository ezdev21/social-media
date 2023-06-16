# The Social Media project


**make sure to configure pusher credentials from pusher website**

PUSHER_APP_ID=

PUSHER_APP_KEY=

PUSHER_APP_SECRET=

PUSHER_APP_CLUSTER=

![social-media](https://user-images.githubusercontent.com/78965149/184528390-ea9d2a96-9d77-4b3a-9a65-54d58e9d40e1.png)

clone the project `git clone https://github.com/ezra02/social-media`

install componser dependencies `composer update`

copy .env files `cp .env.example .env`

generate key `php artisan key:generate`

migrate database `php artisan migrate`

start server `php artisan serve`

## Installation using docker

`docker-compose build social-media`

`docker-compose up -d`

`docker-compose exec social-media php artisan key:generate`

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
