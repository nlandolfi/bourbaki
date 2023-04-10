# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make posterior_event_probabilities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

posterior_event_probabilities.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > posterior_event_probabilities.tex

posterior_event_probabilities.pdf: ../../*.tex ../../trademark.pdf *.tex posterior_event_probabilities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode posterior_event_probabilities.tex
	../../latexrun posterior_event_probabilities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: posterior_event_probabilities.pdf
	open posterior_event_probabilities.pdf

o:
	make open

terms: posterior_event_probabilities.tex
	bbk_sheet -mode ts posterior_event_probabilities.tex

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
	rm -f posterior_event_probabilities.*
	make reset
	bbk_sheet -mode mk > Makefile