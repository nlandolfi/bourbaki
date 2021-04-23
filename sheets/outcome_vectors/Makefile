# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make outcome_vectors.pdf
	make graph.pdf

outcome_vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > outcome_vectors.tex

outcome_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex outcome_vectors.tex
	# pdflatex --file-line-error -interaction=nonstopmode outcome_vectors.tex
	../../latexrun outcome_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: outcome_vectors.pdf
	open outcome_vectors.pdf

o:
	make open

terms: outcome_vectors.tex
	bbk_sheet -mode ts outcome_vectors.tex

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
	rm -f outcome_vectors.*
	make reset
	bbk_sheet -mode mk > Makefile