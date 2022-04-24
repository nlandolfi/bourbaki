# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make unsupervised_learning_algorithms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

unsupervised_learning_algorithms.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > unsupervised_learning_algorithms.tex

unsupervised_learning_algorithms.pdf: ../../*.tex ../../trademark.pdf *.tex unsupervised_learning_algorithms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode unsupervised_learning_algorithms.tex
	../../latexrun unsupervised_learning_algorithms.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: unsupervised_learning_algorithms.pdf
	open unsupervised_learning_algorithms.pdf

o:
	make open

terms: unsupervised_learning_algorithms.tex
	bbk_sheet -mode ts unsupervised_learning_algorithms.tex

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
	rm -f unsupervised_learning_algorithms.*
	make reset
	bbk_sheet -mode mk > Makefile