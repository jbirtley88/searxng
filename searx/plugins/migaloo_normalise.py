# SPDX-License-Identifier: AGPL-3.0-or-later
# pylint: disable=missing-module-docstring,invalid-name

from flask_babel import gettext

name = gettext('Normalise JSON Responses')
description = gettext('Normalises all JSON responses so that the Migaloo app/ui can be dumb as a box of rocks')
default_on = True


def on_result(request, search, result) -> bool:
    # Check if there is a thumbnail url.
    # If not, then put a blank one in
    # This avoids a lot of heavy-lifting at the client side
    if request.user_agent == 'Migaloo-App':
        try:
            thumbnail = result['thumbnail']
        except:
            try:
                result['thumbnail'] = result['thumbnail_src']
            except:
                result['thumbnail'] = 'https://upload.wikimedia.org/wikipedia/commons/c/ca/1x1.png'
    return True

def post_search(request, search):
    return True
