# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make counting_measure_on_finite_set.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

counting_measure_on_finite_set.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > counting_measure_on_finite_set.tex

counting_measure_on_finite_set.pdf: ../../*.tex ../../trademark.pdf *.tex counting_measure_on_finite_set.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode counting_measure_on_finite_set.tex
	../../latexrun counting_measure_on_finite_set.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: counting_measure_on_finite_set.pdf
	open counting_measure_on_finite_set.pdf

o:
	make open

terms: counting_measure_on_finite_set.tex
	bbk_sheet -mode ts counting_measure_on_finite_set.tex

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
	rm -f counting_measure_on_finite_set.*
	make reset
	bbk_sheet -mode mk > Makefile