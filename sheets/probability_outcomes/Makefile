# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probability_outcomes.pdf

probability_outcomes.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probability_outcomes.tex

probability_outcomes.pdf: ../../*.tex ../../trademark.pdf *.tex probability_outcomes.tex
	# pdflatex --file-line-error -interaction=nonstopmode probability_outcomes.tex
	../../latexrun probability_outcomes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability_outcomes.pdf
	open probability_outcomes.pdf

o:
	make open

terms: probability_outcomes.tex
	bbk_sheet -mode ts probability_outcomes.tex

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
	rm -f probability_outcomes.*
	make reset
	bbk_sheet -mode mk > Makefile