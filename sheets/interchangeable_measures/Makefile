# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make interchangeable_measures.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

interchangeable_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > interchangeable_measures.tex

interchangeable_measures.pdf: ../../*.tex ../../trademark.pdf *.tex interchangeable_measures.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode interchangeable_measures.tex
	../../latexrun interchangeable_measures.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: interchangeable_measures.pdf
	open interchangeable_measures.pdf

o:
	make open

terms: interchangeable_measures.tex
	bbk_sheet -mode ts interchangeable_measures.tex

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
	rm -f interchangeable_measures.*
	make reset
	bbk_sheet -mode mk > Makefile