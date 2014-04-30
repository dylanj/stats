Zepto(function($) {
  body = $('body')
  template = $('#template').html()

  render_page = function(data) {
    body.html(Mustache.render(template, {users: data}))
  }

  $.ajax({
    url: '/api.json',
    type: 'GET',
    dataType: 'json',
    success: render_page,
  })
})
