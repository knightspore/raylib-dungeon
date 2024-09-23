#version 330

in vec2 fragTexCoord;
in vec4 fragColor;
in float time;

out vec4 finalColor;

uniform sampler2D tex;

void main() {
    float pulse = 0.5 + 0.5 * sin(time * 3.0);
    vec4 texColor = texture(tex, fragTexCoord);
    vec2 uv = fragTexCoord;
    finalColor = texColor;
}
