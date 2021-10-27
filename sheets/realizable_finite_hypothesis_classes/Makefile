# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make realizable_finite_hypothesis_classes.pdf

realizable_finite_hypothesis_classes.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > realizable_finite_hypothesis_classes.tex

realizable_finite_hypothesis_classes.pdf: ../../*.tex ../../trademark.pdf *.tex realizable_finite_hypothesis_classes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode realizable_finite_hypothesis_classes.tex
	../../latexrun realizable_finite_hypothesis_classes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: realizable_finite_hypothesis_classes.pdf
	open realizable_finite_hypothesis_classes.pdf

o:
	make open

terms: realizable_finite_hypothesis_classes.tex
	bbk_sheet -mode ts realizable_finite_hypothesis_classes.tex

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
	rm -f realizable_finite_hypothesis_classes.*
	make reset
	bbk_sheet -mode mk > Makefile