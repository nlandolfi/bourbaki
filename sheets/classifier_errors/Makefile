# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make classifier_errors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

classifier_errors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > classifier_errors.tex

classifier_errors.pdf: ../../*.tex ../../trademark.pdf *.tex classifier_errors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode classifier_errors.tex
	../../latexrun classifier_errors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: classifier_errors.pdf
	open classifier_errors.pdf

o:
	make open

terms: classifier_errors.tex
	bbk_sheet -mode ts classifier_errors.tex

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
	rm -f classifier_errors.*
	make reset
	bbk_sheet -mode mk > Makefile