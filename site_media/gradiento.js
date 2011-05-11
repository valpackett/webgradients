var lid = 1;
var lbimgr = 'http://myfreeweb.ru/media/mfwjs/lb/images/';
var lbparams = {
    imageBlank: lbimgr + 'lightbox-blank.gif',
    imageLoading: lbimgr + 'lightbox-ico-loading.gif',
    imageBtnClose: lbimgr + 'lightbox-btn-close.gif',
    imageBtnPrev: lbimgr + 'lightbox-btn-prev.gif',
    imageBtnNext: lbimgr + 'lightbox-btn-next.gif',
};
var imag = new Image(); imag.src = '/make?height=150&width=15';
function drawcanv () {
    $('#demos').append('<div class="gradient" title="Example of a 15px wide gradient on HTML5 canvas repeated by JavaScript"><canvas id="canvgr" width="150" height="150"></canvas></div>');
    var cont = document.getElementById('canvgr').getContext('2d');
    var px;
    for (px=0;px<=150;px+=15) {
	cont.drawImage(imag, px, 0);
    }
    return true
}
function swap (one, two) {
    var v_two = $(two).val();
    $(two).val($(one).val());
    $(one).val(v_two);
}
$(window).load(function () {
    window.drawcanv();
    $('body').append('<div id="light" style="display: none;"></div>');
    $('#biggr').append('<div class="advice">Resize me! &rarr;</div>').resizable({ minHeight: 150, maxHeight: 150, minWidth: 50, containment: '#wrapper' });
    $('.swap').show();
    $('#colswap').click(function () {
	window.swap('input[name=start]', 'input[name=end]');
    $('.color').focus().focus(); // for jscolor refreshing
    });
    $('#sizeswap').click(function () {
	window.swap('input[name=width]', 'input[name=height]');
    });
    $('#tryform').submit(function () {
    var width = $('input[name=width]').val(),
	    height = $('input[name=height]').val(),
	    start = $('input[name=start]').val(),
	    end = $('input[name=end]').val(),
	    url = '/make?width=' + width + '&height=' + height + '&start=' + encodeURIComponent(start) + '&end=' + encodeURIComponent(end);
	$('#light').append('<a id="' + window.lid + '" href="' + url + '"><img src="' + url + '" alt="Image" /></a>');
	$('#light a').lightBox(window.lbparams);
	$('#' + window.lid).click();
	window.lid += 1;
	return false;
    });
    if (!Modernizr.inputtypes.color) {
        $('input[type=color]').addClass('color');
        $('body').append('<scr'+'ipt src="/site_media/jscolor.js" type="text/javascript"></script><script type="text/javascript">jscolor.dir=\'/site_media/\';jscolor.init()</script>');
    }
});

