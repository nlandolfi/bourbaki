# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make event_probabilities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

event_probabilities.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > event_probabilities.tex

event_probabilities.pdf: ../../*.tex ../../trademark.pdf *.tex event_probabilities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode event_probabilities.tex
	../../latexrun event_probabilities.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: event_probabilities.pdf
	open event_probabilities.pdf

o:
	make open

terms: event_probabilities.tex
	bbk sheet -mode ts event_probabilities.tex

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
	rm -f event_probabilities.*
	make reset
	bbk sheet -mode mk > Makefile