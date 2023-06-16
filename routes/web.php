<?php

use App\Http\Controllers\ChatController;
use App\Http\Controllers\FriendController;
use App\Http\Controllers\GroupController;
use Illuminate\Support\Facades\Route;
use Illuminate\Support\Facades\Auth;
use App\Http\Controllers\HomeController;
use App\Http\Controllers\PostController;
use App\Http\Controllers\StoryController;
use App\Http\Controllers\UserController;
use App\Http\Controllers\VideoController;
use App\Models\Message;

Route::get('/',[HomeController::class,'index']);

Auth::routes();

Route::get('/home',[HomeController::class, 'index'])->name('home');

Route::get('/auth/{provider}/redirect',[SocialiteController::class,'redirect']);
Route::get('/auth/{provider}/callback',[SocialiteController::class,'callback']);


Route::prefix('post')->group(function(){
  Route::get('create',[PostController::class,'create']);
  Route::post('store',[PostController::class,'store'])->name('post.store');
  Route::get('edit/{id}',[PostController::class,'edit']);
  Route::patch('update',[PostController::class,'update'])->name('post.update');
  Route::delete('delete',[PostController::class,'delete']);
});

Route::prefix('story')->group(function(){
  Route::get('create',[StoryController::class,'create']);
  Route::post('store',[StoryController::class,'store']);
  Route::get('edit/{id}',[StryController::class,'edit']);
  Route::get('delete',[StoryController::class,'delete']);
});

Route::prefix('video')->group(function(){
  Route::get('',[VideoController::class,'index']);
  Route::get('create',[VideoController::class,'create']);
  Route::post('store',[VideoController::class,'store']);
  Route::get('edit/{id}',[VideoController::class,'edit']);
  Route::get('delete',[VideoController::class,'delete']);
});

Route::prefix('friend')->group(function(){
  Route::get('/',[FriendController::class,'index']);
  Route::get('friends',[FriendController::class,'friends']);
  Route::post('follow',[FriendController::class,'follow']);
  Route::get('unfollow',[FriendController::class,'unFollow']);
  Route::get('block',[FriendController::class,'block']);
});

Route::prefix('group')->group(function(){
    Route::get('create',[GroupController::class,'create']);
    Route::post('store',[GroupController::class,'store']);
    Route::get('edit/{id}',[GroupController::class,'edit']);
    Route::get('delete',[GroupController::class,'delete']);
    Route::prefix('member')->group(function(){
      Route::get('all',[GroupController::class,'members']);
      Route::post('add',[GroupController::class,'addmemember']);
      Route::post('remove',[GroupController::class,'removeMember']);
    });
  });

  Route::prefix('channel')->group(function(){
    Route::get('create',[ChannelController::class,'create']);
    Route::post('store',[ChannelController::class,'store']);
    Route::get('edit/{id}',[ChannelController::class,'edit']);
    Route::get('delete',[ChannelController::class,'delete']);
  });

  Route::prefix('comment')->group(function(){
    Route::get('create',[CommentController::class,'create']);
    Route::post('store',[CommentController::class,'store']);
    Route::get('edit/{id}',[CommentController::class,'edit']);
    Route::get('delete',[CommentController::class,'delete']);
  });

  Route::prefix('profile')->group(function(){
    Route::get('create',[ProfileController::class,'create']);
    Route::post('store',[ProfileController::class,'store']);
    Route::get('edit/{id}',[ProfileController::class,'edit']);
    Route::get('delete',[ProfileController::class,'delete']);
  });

  Route::prefix('chat')->group(function(){
    Route::get('/',[ChatController::class,'index'])->name('chat');
    Route::get('friends',[ChatController::class,'friends']);
    Route::get('messages',[ChatController::class,'messages']);
    Route::post('message/send',[ChatController::class,'sendMessage']);
    Route::post('store',[ChatController::class,'store']);
    Route::get('message/edit',[ChatController::class,'edit']);
    Route::get('message/delete',[ChatController::class,'delete']);
  });

  Route::prefix('user')->group(function(){
   Route::get('notifications',[UserController::class,'notifications']);
   Route::post('notification/mark-as-read',[UserController::class,'markAsRead']);
  });