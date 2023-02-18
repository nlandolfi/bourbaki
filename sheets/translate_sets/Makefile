# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make translate_sets.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

translate_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > translate_sets.tex

translate_sets.pdf: ../../*.tex ../../trademark.pdf *.tex translate_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode translate_sets.tex
	../../latexrun translate_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: translate_sets.pdf
	open translate_sets.pdf

o:
	make open

terms: translate_sets.tex
	bbk_sheet -mode ts translate_sets.tex

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
	rm -f translate_sets.*
	make reset
	bbk_sheet -mode mk > Makefile