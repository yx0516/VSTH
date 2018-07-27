var code;

function createCode() {
    code = "";
    var codeLength = 4;
    var checkCode = document.getElementById("checkCode");
    var codeChars = new Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
        'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
        'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z');
    for (var i = 0; i < codeLength; i++) {
        var charNum = Math.floor(Math.random() * 52);
        code += codeChars[charNum];
    }
    if (checkCode) {
        checkCode.className = "authenticode";
        checkCode.innerHTML = code;
    }
}

function validateCode() {
    var inputCode = document.getElementById("form-code").value;
    if (inputCode.length <= 0) {
        alert("Please enter Verification Code！");
        return false;
    } else if (inputCode.toUpperCase() != code.toUpperCase()) {
        alert("Verification Code is wrong！");
        createCode();
        return false;
    } else {
        return true;
    }
}
//邮箱验证
function isEmail(str) {
    var reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/;
    return reg.test(str);
}

$(document).ready(function() {
    createCode();
    //Form validation
    $('.login-form input[type="text"], .login-form input[type="password"], .login-form textarea').on('focus', function() {
        $(this).removeClass('input-error');
    });

    $("#form-useremail").change(function() {
        if (!isEmail($(this).val())) {
            alert("Please enter a valid email address !");
            return false;
        }
    });
    $("#form-password-2").change(function() {
        if ($(this).val() != $("#form-password").val()) {
            alert("Password is not right!");
            return false;
        }
    });
    $("#checkCode").on("click", function() {
        createCode();
    });
    $("#refreshCode").on("click", function() {
        createCode();
    });

    //login
    $('.submitInfo').on('click', function(e) {
        $(".form-group").find('input[type="text"], input[type="password"]').each(function() {
            if ($(this).val() == "") {
                e.preventDefault();
                $(this).addClass('input-error');
            } else {
                $(this).removeClass('input-error');
            }
        });
        if ($("input[type='text']").hasClass('input-error') && $("input[type='password']").hasClass('input-error')) {
            e.preventDefault();
            return false;
        } else {
            var username = $("#form-username").val();
            var userpwd = $("#form-password").val();
            var remember = "";
            if ($("input[type='checkbox']").is(':checked')) {
                remember = "yes";
            } else {
                remember = "no";
            }
            $.ajax({
                url: "/account/login",
                type: "POST",
                data: {
                    username: username,
                    password: userpwd,
                    remember: remember
                },
                dataType: "html",
            }).done(loginRedirct);
        }
    });

    $('.signBtn').on('click', function(e) {
        $(".form-group").find('input[type="text"], input[type="password"]').each(function() {
            if ($(this).val() == "") {
                e.preventDefault();
                $(this).addClass('input-error');
            } else {
                $(this).removeClass('input-error');
            }
        });

        if ($("input[type='text']").hasClass('input-error') && $("input[type='password']").hasClass('input-error')) {
            //alert("both empty");
            e.preventDefault();
            return false;
        } else {
            if (validateCode()) {
                var username = $("#form-username").val();
                var userpwd = $("#form-password").val();
                var email = $('#form-useremail').val();
                $.ajax({
                    url: "/account/register",
                    type: "POST",
                    data: {
                        username: username,
                        password: userpwd,
                        email: email
                    },
                    dataType: "html",
                }).done(Redirct);
            }
        }
    });




    $('.findpwdBtn').on('click', function(e) {
        $(".form-group").find('input[type="text"], input[type="password"]').each(function() {
            if ($(this).val() == "") {
                e.preventDefault();
                $(this).addClass('input-error');
            } else {
                $(this).removeClass('input-error');
            }
        });
        if ($("input[type='text']").hasClass('input-error') && $("input[type='password']").hasClass('input-error')) {
            //alert("both empty");
            e.preventDefault();
            return
        } else {
            var username = $("#form-username").val();
            var email = $('#form-useremail').val();
            // var userpwd = $("#form-password").val();

            // var validation = $('#form-code').val();
            $.ajax({
                url: "/account/GetPassword",
                type: "POST",
                data: {
                    username: username,
                    email: email
                },
                dataType: "html",
            }).done(Redirct);
        }
    });
});

// signup findPwd 成功回调
function Redirct(msg) {
    var msgJson = JSON.parse(msg);
    console.log(msgJson);
    if (msgJson.State == -1) {
        alert(msgJson.Msg);
        return false;
    } else {
        alert(msgJson.Msg);
        window.location.href = 'login.html';
    }
}

// login成功回调
function loginRedirct(msg) {
    var msgJson = JSON.parse(msg);
    console.log(msgJson);
    if (msgJson.State == -1) {
        alert(msgJson.Msg);
        return false;
    } else {
        // @todo : 存储用户的登录信息
        setLocalStorage(msgJson.Data);
        // @todo : 登录之后跳转到首页，带有用户名
        window.location.href = 'index.html';
    }
}

function setLocalStorage(userInfo) {
    //存储，IE6~7 cookie 其他浏览器HTML5本地存储
    if (window.localStorage) {
        localStorage.setItem("userid", userInfo.userid);
        localStorage.setItem("userName", userInfo.username);
    } else {
        Cookie.write("userid", userInfo.userid);
        Cookie.write("userName", userInfo.username);
    }
}