function isLog(msg) {
    var jobInfo = JSON.parse(msg);
    console.log(jobInfo);

    var log_ = $('#log');
    if (jobInfo.State == -1) {
        log_.removeClass('logOut');
        log_.addClass('logIn');
        log_.html('Log in');
        log_.attr(href, '/login.html');
        $(".user-info span").html();
        alert(jobInfo.Msg);
        return false;
    } else {
        if (jobInfo.Data.userid < 0) {
            log_.removeClass('logOut');
            log_.addClass('logIn');
            log_.html('Log in');
            log_.attr(href, '/login.html');
            $(".user-info span").html();
        } else {
            log_.removeClass('logIn');
            log_.addClass('logOut');
            log_.html('Log out');
            log_.attr(href, '/account/logout');
            $(".user-info span").html(jobInfo.Data.username);
        }
    }
}