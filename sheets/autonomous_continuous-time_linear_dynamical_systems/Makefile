# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make autonomous_continuous-time_linear_dynamical_systems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

autonomous_continuous-time_linear_dynamical_systems.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > autonomous_continuous-time_linear_dynamical_systems.tex

autonomous_continuous-time_linear_dynamical_systems.pdf: ../../*.tex ../../trademark.pdf *.tex autonomous_continuous-time_linear_dynamical_systems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode autonomous_continuous-time_linear_dynamical_systems.tex
	../../latexrun autonomous_continuous-time_linear_dynamical_systems.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: autonomous_continuous-time_linear_dynamical_systems.pdf
	open autonomous_continuous-time_linear_dynamical_systems.pdf

o:
	make open

terms: autonomous_continuous-time_linear_dynamical_systems.tex
	bbk sheet -mode ts autonomous_continuous-time_linear_dynamical_systems.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f autonomous_continuous-time_linear_dynamical_systems.*
	make reset
	bbk sheet -mode mk > Makefile