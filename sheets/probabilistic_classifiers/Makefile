# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probabilistic_classifiers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

probabilistic_classifiers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probabilistic_classifiers.tex

probabilistic_classifiers.pdf: ../../*.tex ../../trademark.pdf *.tex probabilistic_classifiers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode probabilistic_classifiers.tex
	../../latexrun probabilistic_classifiers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probabilistic_classifiers.pdf
	open probabilistic_classifiers.pdf

o:
	make open

terms: probabilistic_classifiers.tex
	bbk_sheet -mode ts probabilistic_classifiers.tex

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
	rm -f probabilistic_classifiers.*
	make reset
	bbk_sheet -mode mk > Makefile