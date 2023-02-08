import cocos

from layers import MouseDisplay

from comm import connect, send_to_server, start_message_queue_thread, start_receiver_thread
from messages import PlayerLogin


def main():

    connect()

    # Start the background threads needed before c
    mq_thread = start_message_queue_thread()

    rec_thread = start_receiver_thread()

    # Login
    send_to_server(PlayerLogin("player", "s3cret"))

    # Init director
    cocos.director.director.init()

    mouse_display = MouseDisplay()

    main_scene = cocos.scene.Scene(mouse_display)

    # And let the director run the scene
    cocos.director.director.run(main_scene)

    mq_thread.join()
    rec_thread.join()


if __name__ == "__main__":
    main()
