# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make standard_deviation.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

standard_deviation.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > standard_deviation.tex

standard_deviation.pdf: ../../*.tex ../../trademark.pdf *.tex standard_deviation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode standard_deviation.tex
	../../latexrun standard_deviation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: standard_deviation.pdf
	open standard_deviation.pdf

o:
	make open

terms: standard_deviation.tex
	bbk_sheet -mode ts standard_deviation.tex

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
	rm -f standard_deviation.*
	make reset
	bbk_sheet -mode mk > Makefile